#coding=utf-8
import unittest

class TestCaseMySolution(unittest.TestCase):
    def testCase01(self):
        nums = [2,7,11,15]
        target= 9
        want = [0,1]

        got = Solution().twoSum(nums, target)
        self.assertEqual(want,got)

    def testCase02(self):
        nums = [3,2,4]
        target= 6
        want = [1,2]

        got = Solution().twoSum(nums, target)
        self.assertEqual(want,got)

        # self.assertEqualTure('FOO'.isupper())
        # self.assertFalse('Foo'.isupper())
        # self.assertEqual('foo',upper(),'FOO')
        # self.assertEqual('hello world'.split(),['hello','world']) with self.assertEqualRaise(TypeError):

class Solution:
    def twoSum(self, nums, target) :
        """
        :type nums: List[int]
        :type target: int
        :rtype: List[int]
        """
        for i in range(0,len(nums)-1) :
            for j in range(i+1,len(nums)) :
                if nums[i]+nums[j] == target :
                    return [i,j]
        return []


if __name__ == '__main__':
    suite = unittest.TestSuite()
    suite.addTest(unittest.TestLoader().loadTestsFromTestCase(TestCaseMySolution))
    testResult=unittest.TextTestRunner(verbosity=2).run(suite)

    print("testsRun:%s" % testResult.testsRun)
    print("failures:%s" % len(testResult.failures))
    print("errors  :%s" % len(testResult.errors))
    print("skipped :%s" % len(testResult.skipped))
