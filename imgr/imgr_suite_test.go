package imgr_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestImgr(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Imgr Suite")
}
