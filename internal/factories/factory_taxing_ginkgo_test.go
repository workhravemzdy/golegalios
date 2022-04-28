package factories

import (
	"github.com/hravemzdy/golegalios/internal/types"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Factory Taxing For Year 2009-2022", func() {
	DescribeTable("Factory Taxing Samples",
		func(tt testEntry) {
			factory := NewFactoryTaxing()
			period := types.GetPeriodWithYearMonth(tt.testYear, tt.testMonth)
			props, exists := factory.GetProps(period)

			Expect(exists).To(BeTrue())
			Expect(props).ToNot(BeNil())
			var resultYear = props.GetVersionValue()
			Expect(resultYear).To(Equal(tt.resultYear))
		},
		Entry("2009-1", testEntry{2009, 1, 0}),
		Entry("2009-2", testEntry{2009, 2, 0}),
		Entry("2009-3", testEntry{2009, 3, 0}),
		Entry("2009-4", testEntry{2009, 4, 0}),
		Entry("2009-5", testEntry{2009, 5, 0}),
		Entry("2009-6", testEntry{2009, 6, 0}),
		Entry("2009-7", testEntry{2009, 7, 0}),
		Entry("2009-8", testEntry{2009, 8, 0}),
		Entry("2009-9", testEntry{2009, 9, 0}),
		Entry("2009-10", testEntry{2009, 10, 0}),
		Entry("2009-11", testEntry{2009, 11, 0}),
		Entry("2009-12", testEntry{2009, 12, 0}),
		Entry("2010-1", testEntry{2010, 1, 2010}),
		Entry("2010-2", testEntry{2010, 2, 2010}),
		Entry("2010-3", testEntry{2010, 3, 2010}),
		Entry("2010-4", testEntry{2010, 4, 2010}),
		Entry("2010-5", testEntry{2010, 5, 2010}),
		Entry("2010-6", testEntry{2010, 6, 2010}),
		Entry("2010-7", testEntry{2010, 7, 2010}),
		Entry("2010-8", testEntry{2010, 8, 2010}),
		Entry("2010-9", testEntry{2010, 9, 2010}),
		Entry("2010-10", testEntry{2010, 10, 2010}),
		Entry("2010-11", testEntry{2010, 11, 2010}),
		Entry("2010-12", testEntry{2010, 12, 2010}),
		Entry("2011-1", testEntry{2011, 1, 2011}),
		Entry("2011-2", testEntry{2011, 2, 2011}),
		Entry("2011-3", testEntry{2011, 3, 2011}),
		Entry("2011-4", testEntry{2011, 4, 2011}),
		Entry("2011-5", testEntry{2011, 5, 2011}),
		Entry("2011-6", testEntry{2011, 6, 2011}),
		Entry("2011-7", testEntry{2011, 7, 2011}),
		Entry("2011-8", testEntry{2011, 8, 2011}),
		Entry("2011-9", testEntry{2011, 9, 2011}),
		Entry("2011-10", testEntry{2011, 10, 2011}),
		Entry("2011-11", testEntry{2011, 11, 2011}),
		Entry("2011-12", testEntry{2011, 12, 2011}),
		Entry("2012-1", testEntry{2012, 1, 2012}),
		Entry("2012-2", testEntry{2012, 2, 2012}),
		Entry("2012-3", testEntry{2012, 3, 2012}),
		Entry("2012-4", testEntry{2012, 4, 2012}),
		Entry("2012-5", testEntry{2012, 5, 2012}),
		Entry("2012-6", testEntry{2012, 6, 2012}),
		Entry("2012-7", testEntry{2012, 7, 2012}),
		Entry("2012-8", testEntry{2012, 8, 2012}),
		Entry("2012-9", testEntry{2012, 9, 2012}),
		Entry("2012-10", testEntry{2012, 10, 2012}),
		Entry("2012-11", testEntry{2012, 11, 2012}),
		Entry("2012-12", testEntry{2012, 12, 2012}),
		Entry("2013-1", testEntry{2013, 1, 2013}),
		Entry("2013-2", testEntry{2013, 2, 2013}),
		Entry("2013-3", testEntry{2013, 3, 2013}),
		Entry("2013-4", testEntry{2013, 4, 2013}),
		Entry("2013-5", testEntry{2013, 5, 2013}),
		Entry("2013-6", testEntry{2013, 6, 2013}),
		Entry("2013-7", testEntry{2013, 7, 2013}),
		Entry("2013-8", testEntry{2013, 8, 2013}),
		Entry("2013-9", testEntry{2013, 9, 2013}),
		Entry("2013-10", testEntry{2013, 10, 2013}),
		Entry("2013-11", testEntry{2013, 11, 2013}),
		Entry("2013-12", testEntry{2013, 12, 2013}),
		Entry("2014-1", testEntry{2014, 1, 2014}),
		Entry("2014-2", testEntry{2014, 2, 2014}),
		Entry("2014-3", testEntry{2014, 3, 2014}),
		Entry("2014-4", testEntry{2014, 4, 2014}),
		Entry("2014-5", testEntry{2014, 5, 2014}),
		Entry("2014-6", testEntry{2014, 6, 2014}),
		Entry("2014-7", testEntry{2014, 7, 2014}),
		Entry("2014-8", testEntry{2014, 8, 2014}),
		Entry("2014-9", testEntry{2014, 9, 2014}),
		Entry("2014-10", testEntry{2014, 10, 2014}),
		Entry("2014-11", testEntry{2014, 11, 2014}),
		Entry("2014-12", testEntry{2014, 12, 2014}),
		Entry("2015-1", testEntry{2015, 1, 2015}),
		Entry("2015-2", testEntry{2015, 2, 2015}),
		Entry("2015-3", testEntry{2015, 3, 2015}),
		Entry("2015-4", testEntry{2015, 4, 2015}),
		Entry("2015-5", testEntry{2015, 5, 2015}),
		Entry("2015-6", testEntry{2015, 6, 2015}),
		Entry("2015-7", testEntry{2015, 7, 2015}),
		Entry("2015-8", testEntry{2015, 8, 2015}),
		Entry("2015-9", testEntry{2015, 9, 2015}),
		Entry("2015-10", testEntry{2015, 10, 2015}),
		Entry("2015-11", testEntry{2015, 11, 2015}),
		Entry("2015-12", testEntry{2015, 12, 2015}),
		Entry("2016-1", testEntry{2016, 1, 2016}),
		Entry("2016-2", testEntry{2016, 2, 2016}),
		Entry("2016-3", testEntry{2016, 3, 2016}),
		Entry("2016-4", testEntry{2016, 4, 2016}),
		Entry("2016-5", testEntry{2016, 5, 2016}),
		Entry("2016-6", testEntry{2016, 6, 2016}),
		Entry("2016-7", testEntry{2016, 7, 2016}),
		Entry("2016-8", testEntry{2016, 8, 2016}),
		Entry("2016-9", testEntry{2016, 9, 2016}),
		Entry("2016-10", testEntry{2016, 10, 2016}),
		Entry("2016-11", testEntry{2016, 11, 2016}),
		Entry("2016-12", testEntry{2016, 12, 2016}),
		Entry("2017-1", testEntry{2017, 1, 2017}),
		Entry("2017-2", testEntry{2017, 2, 2017}),
		Entry("2017-3", testEntry{2017, 3, 2017}),
		Entry("2017-4", testEntry{2017, 4, 2017}),
		Entry("2017-5", testEntry{2017, 5, 2017}),
		Entry("2017-6", testEntry{2017, 6, 2017}),
		Entry("2017-7", testEntry{2017, 7, 2017}),
		Entry("2017-8", testEntry{2017, 8, 2017}),
		Entry("2017-9", testEntry{2017, 9, 2017}),
		Entry("2017-10", testEntry{2017, 10, 2017}),
		Entry("2017-11", testEntry{2017, 11, 2017}),
		Entry("2017-12", testEntry{2017, 12, 2017}),
		Entry("2018-1", testEntry{2018, 1, 2018}),
		Entry("2018-2", testEntry{2018, 2, 2018}),
		Entry("2018-3", testEntry{2018, 3, 2018}),
		Entry("2018-4", testEntry{2018, 4, 2018}),
		Entry("2018-5", testEntry{2018, 5, 2018}),
		Entry("2018-6", testEntry{2018, 6, 2018}),
		Entry("2018-7", testEntry{2018, 7, 2018}),
		Entry("2018-8", testEntry{2018, 8, 2018}),
		Entry("2018-9", testEntry{2018, 9, 2018}),
		Entry("2018-10", testEntry{2018, 10, 2018}),
		Entry("2018-11", testEntry{2018, 11, 2018}),
		Entry("2018-12", testEntry{2018, 12, 2018}),
		Entry("2019-1", testEntry{2019, 1, 2019}),
		Entry("2019-2", testEntry{2019, 2, 2019}),
		Entry("2019-3", testEntry{2019, 3, 2019}),
		Entry("2019-4", testEntry{2019, 4, 2019}),
		Entry("2019-5", testEntry{2019, 5, 2019}),
		Entry("2019-6", testEntry{2019, 6, 2019}),
		Entry("2019-7", testEntry{2019, 7, 2019}),
		Entry("2019-8", testEntry{2019, 8, 2019}),
		Entry("2019-9", testEntry{2019, 9, 2019}),
		Entry("2019-10", testEntry{2019, 10, 2019}),
		Entry("2019-11", testEntry{2019, 11, 2019}),
		Entry("2019-12", testEntry{2019, 12, 2019}),
		Entry("2020-1", testEntry{2020, 1, 2020}),
		Entry("2020-2", testEntry{2020, 2, 2020}),
		Entry("2020-3", testEntry{2020, 3, 2020}),
		Entry("2020-4", testEntry{2020, 4, 2020}),
		Entry("2020-5", testEntry{2020, 5, 2020}),
		Entry("2020-6", testEntry{2020, 6, 2020}),
		Entry("2020-7", testEntry{2020, 7, 2020}),
		Entry("2020-8", testEntry{2020, 8, 2020}),
		Entry("2020-9", testEntry{2020, 9, 2020}),
		Entry("2020-10", testEntry{2020, 10, 2020}),
		Entry("2020-11", testEntry{2020, 11, 2020}),
		Entry("2020-12", testEntry{2020, 12, 2020}),
		Entry("2021-1", testEntry{2021, 1, 2021}),
		Entry("2021-2", testEntry{2021, 2, 2021}),
		Entry("2021-3", testEntry{2021, 3, 2021}),
		Entry("2021-4", testEntry{2021, 4, 2021}),
		Entry("2021-5", testEntry{2021, 5, 2021}),
		Entry("2021-6", testEntry{2021, 6, 2021}),
		Entry("2021-7", testEntry{2021, 7, 2021}),
		Entry("2021-8", testEntry{2021, 8, 2021}),
		Entry("2021-9", testEntry{2021, 9, 2021}),
		Entry("2021-10", testEntry{2021, 10, 2021}),
		Entry("2021-11", testEntry{2021, 11, 2021}),
		Entry("2021-12", testEntry{2021, 12, 2021}),
		Entry("2022-1", testEntry{2022, 1, 2022}),
		Entry("2022-2", testEntry{2022, 2, 2022}),
		Entry("2022-3", testEntry{2022, 3, 2022}),
		Entry("2022-4", testEntry{2022, 4, 2022}),
		Entry("2022-5", testEntry{2022, 5, 2022}),
		Entry("2022-6", testEntry{2022, 6, 2022}),
		Entry("2022-7", testEntry{2022, 7, 2022}),
		Entry("2022-8", testEntry{2022, 8, 2022}),
		Entry("2022-9", testEntry{2022, 9, 2022}),
		Entry("2022-10", testEntry{2022, 10, 2022}),
		Entry("2022-11", testEntry{2022, 11, 2022}),
		Entry("2022-12", testEntry{2022, 12, 2022}),
	)
})
