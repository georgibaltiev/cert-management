/*
 * // SPDX-FileCopyrightText: SAP SE or an SAP affiliate company and Gardener contributors
 * //
 * // SPDX-License-Identifier: Apache-2.0
 */

package issuer_test

import (
	"context"

	"github.com/gardener/controller-manager-library/pkg/controllermanager"
	"github.com/gardener/controller-manager-library/pkg/ctxutil"
	. "github.com/gardener/gardener/pkg/utils/test/matchers"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/utils/ptr"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/gardener/cert-management/pkg/apis/cert/v1alpha1"
)

var _ = Describe("Issuer controller tests", func() {
	var (
		testRunID     string
		testNamespace *corev1.Namespace
	)

	BeforeEach(func() {
		Expect(acmeDirectoryAddress).NotTo(BeEmpty())

		ctxLocal := context.Background()
		ctx0 := ctxutil.CancelContext(ctxutil.WaitGroupContext(context.Background(), "main"))
		ctx = ctxutil.TickContext(ctx0, controllermanager.DeletionActivity)

		By("Create test Namespace")
		testNamespace = &corev1.Namespace{
			ObjectMeta: metav1.ObjectMeta{
				GenerateName: "issuer-",
			},
		}
		Expect(testClient.Create(ctxLocal, testNamespace)).To(Succeed())
		log.Info("Created Namespace for test", "namespaceName", testNamespace.Name)
		testRunID = testNamespace.Name

		DeferCleanup(func() {
			By("Delete test Namespace")
			Expect(testClient.Delete(ctxLocal, testNamespace)).To(Or(Succeed(), BeNotFoundError()))
		})

		By("Start manager")

		go func() {
			defer GinkgoRecover()
			args := []string{
				"--kubeconfig", kubeconfigFile,
				"--controllers", "issuer",
				"--issuer-namespace", testRunID,
				"--omit-lease",
				"--pool.size", "1",
			}
			runControllerManager(ctx, args)
		}()

		DeferCleanup(func() {
			By("Stop manager")
			if ctx != nil {
				ctxutil.Cancel(ctx)
			}
		})
	})

	Context("ACME issuer", func() {
		It("should create an ACME issuer", func() {
			issuer := &v1alpha1.Issuer{
				ObjectMeta: metav1.ObjectMeta{
					Namespace: testRunID,
					Name:      "acme1",
				},
				Spec: v1alpha1.IssuerSpec{
					ACME: &v1alpha1.ACMESpec{
						Email:            "foo@somewhere-foo-123456.com",
						Server:           acmeDirectoryAddress,
						AutoRegistration: true,
					},
				},
			}
			Expect(testClient.Create(ctx, issuer)).To(Succeed())
			DeferCleanup(func() {
				Expect(testClient.Delete(ctx, issuer)).To(Succeed())
			})

			Eventually(func(g Gomega) {
				Expect(testClient.Get(ctx, client.ObjectKeyFromObject(issuer), issuer)).To(Succeed())
				g.Expect(issuer.Status.State).To(Equal("Ready"))
			}).Should(Succeed())
		})
	})

	Context("Self-signed issuer", func() {
		It("should be able to create self-signed certificates", func() {
			By("Create self-signed issuer")
			issuer := &v1alpha1.Issuer{
				ObjectMeta: metav1.ObjectMeta{
					Namespace: testRunID,
					Name:      "self-signed-issuer",
				},
				Spec: v1alpha1.IssuerSpec{
					SelfSigned: &v1alpha1.SelfSignedSpec{},
				},
			}
			Expect(testClient.Create(ctx, issuer)).To(Succeed())
			DeferCleanup(func() {
				Expect(testClient.Delete(ctx, issuer)).To(Succeed())
			})

			Eventually(func(g Gomega) {
				Expect(testClient.Get(ctx, client.ObjectKeyFromObject(issuer), issuer)).To(Succeed())
				g.Expect(issuer.Status.State).To(Equal("Ready"))
			}).Should(Succeed())

			By("Create self-signed certificate")
			certificate := &v1alpha1.Certificate{
				ObjectMeta: metav1.ObjectMeta{
					Namespace: testRunID,
					Name:      "self-signed-certificate",
				},
				Spec: v1alpha1.CertificateSpec{
					CommonName: ptr.To("ca1.mydomain.com"),
					IsCA:       ptr.To(true),
					IssuerRef: &v1alpha1.IssuerRef{
						Name:      issuer.Name,
						Namespace: issuer.Namespace,
					},
				},
			}
			Expect(testClient.Create(ctx, certificate)).To(Succeed())
			DeferCleanup(func() {
				Expect(testClient.Delete(ctx, certificate)).To(Succeed())
			})

			Eventually(func(g Gomega) {
				Expect(testClient.Get(ctx, client.ObjectKeyFromObject(certificate), certificate)).To(Succeed())
				g.Expect(certificate.Status.State).To(Equal("Ready"))
			}).Should(Succeed())

			By("Resolve certificate secret reference")
			secretReference := certificate.Spec.SecretRef
			secretKey := client.ObjectKey{Name: secretReference.Name, Namespace: secretReference.Namespace}
			secret := &corev1.Secret{}
			Expect(testClient.Get(ctx, secretKey, secret)).To(Succeed())
			Expect(secret.Data).To(HaveKeyWithValue("ca.crt", Not(BeEmpty())))
			Expect(secret.Data).To(HaveKeyWithValue("tls.crt", Not(BeEmpty())))
			Expect(secret.Data).To(HaveKeyWithValue("tls.key", Not(BeEmpty())))
		})
	})
})
