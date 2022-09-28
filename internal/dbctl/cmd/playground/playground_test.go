/*
Copyright © 2022 The dbctl Authors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package playground

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"k8s.io/cli-runtime/pkg/genericclioptions"
)

var _ = Describe("playground", func() {
	var streams genericclioptions.IOStreams

	BeforeEach(func() {
		streams, _, _, _ = genericclioptions.NewTestIOStreams()
	})

	It("New playground command", func() {
		cmd := NewPlaygroundCmd(streams)
		Expect(cmd != nil).Should(BeTrue())
		Expect(cmd.HasSubCommands()).Should(BeTrue())
	})

	It("init command", func() {
		cmd := newInitCmd(streams)
		Expect(cmd != nil).Should(BeTrue())

		o := &initOptions{
			Replicas: 0,
		}
		Expect(o.validate()).To(MatchError("replicas should greater than 0"))

		o.Replicas = 1
		Expect(o.validate()).Should(Succeed())
		Expect(o.run()).Should(HaveOccurred())
	})

	It("destroy command", func() {
		cmd := newDestroyCmd(streams)
		Expect(cmd != nil).Should(BeTrue())

		o := &destroyOptions{
			IOStreams: streams,
		}
		Expect(o.destroyPlayground()).Should(HaveOccurred())
	})
})
