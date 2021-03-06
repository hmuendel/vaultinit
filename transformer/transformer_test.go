package transformer_test

import (
	"github.com/hmuendel/kubevaulter/transformer"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Transformer", func() {

	var fnMap = transformer.DefaultFuncMap()

	Describe("Applying the Hash function directly", func() {
		Context("for the sha1", func() {
			It("should create the right hash value", func() {
				Expect(transformer.Sha1("test")).To(Equal(
					"a94a8fe5ccb19ba61c4c0873d391e987982fbbd3"))
			})
		})
		Context("for the sha256", func() {
			It("should create the right hash value", func() {
				Expect(transformer.Sha256("test")).To(Equal(
					"9f86d081884c7d659a2feaa0c55ad015a3bf4f1b2b0b822cd15d6c15b0f00a08"))
			})
		})
	})
	Describe("Applying the hash function from the default map", func() {
		Context("for the sha256", func() {
			It("shoud create the right hash value", func() {
				Expect(fnMap["SHA256"]("test")).To(Equal(
					"9f86d081884c7d659a2feaa0c55ad015a3bf4f1b2b0b822cd15d6c15b0f00a08"))
			})
		})
	})

})
