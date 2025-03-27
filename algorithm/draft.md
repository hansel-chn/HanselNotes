# 字符串匹配算法

* kmp, bm, sunday algorithm

# Dp

* 最长递增子序列构建Dp数组时，采用包含最末尾的数构成的最长子序列，因为如果不包含，dp数组扩展时，不能确定是否应当扩展，不知道新添加的数是否基于之前数组递增，比如：
  [4，5，9，6，7]，dp[9]=dp[6]，7比6大不知道7是否仍然是递增序列

# 关于前序遍历和回溯想到的

* 遍历树更像是已经确定了每个结点的结果，而回溯更倾向于选择自己的路径，比如5个数选3个数组合，回溯更像是通过自己的选择，
  构建出类似树的结构，每一行对应一个盒子，装不同的球。所有遍历是遍历已有结构，获取节点的相关信息，对于多叉树而言，更关心遍历的顺序，
  所以处理数据（比如输出节点的值在循环的外部-前序遍历和后序遍历）；对于回溯而言，大部分情况未关注顺序，一般都是先序处理的？，
  更将倾向于选择，以便在之后，检测经过当前选择后得到的结果，是否满足期望的结果。所以在循环内针对每次选择处理当前数据，而非循环外。
* 有点感觉是回溯的选择，实时生成了“树的结点的状态”，然后进行的是类似dfs对结点的处理
* 后序遍历和分解问题（分冶法）感觉殊途同归

# Binary search

https://labuladong.online/algo/essential-technique/binary-search-framework-2/#%E4%B8%89%E3%80%81%E5%AF%BB%E6%89%BE%E5%8F%B3%E4%BE%A7%E8%BE%B9%E7%95%8C%E7%9A%84%E4%BA%8C%E5%88%86%E6%9F%A5%E6%89%BE

* Supplement
    *
  讲了很多种变形，核心点感觉应该是由于获取mid的index所导致的，不管如何选取子segment，最终都可以通过特殊的处理方法达到目的，文中给出的模板选取的方法最后都不需要特殊处理，其实针对mid的index的位置的特点在取子segment时，采取了相对应的处理办法，也导致了不同模板中left和right取值不一样。
    * 比如为什么左闭右开区间和左闭右闭区间right选择不同赋值方式`right = mid or right = mid - 1`
    * 包括探讨循环条件，为何`left < right or left <= rigth`
        * 循环不能终止和最后一次循环是否进行的问题

# Edit Distance

https://leetcode.com/problems/edit-distance/solutions/25846/c-o-n-space-dp/

# Debate

For the general case to convert`word1[0..i)`to`word2[0..j)`, we break this problem down into sub-problems. Suppose we
have already known how to convert`word1[0..i - 1)`to`word2[0..j - 1)`(`dp[i - 1][j - 1]`), if
`word1[i - 1] == word2[j - 1]`, then no more operation is needed and`dp[i][j] = dp[i - 1][j - 1]`.

why `word[i - 1]`and`word2[j - 1]`are equal,`dp[i][j]`can directly set to`dp[i - 1][j - 1]`rather than
`dp[i][j] = min(dp[i - 1][j - 1], dp[i - 1][j] + 1, dp[i][j - 1] + 1)`.

考虑`dp[i - 1][j - 1], dp[i][j - 1] + 1`,

* 如果`dp[i - 1][j - 1] < dp[i][j - 1]`,`dp[i - 1][j - 1] < dp[i][j - 1] + 1`天然成立
* 如果`dp[i - 1][j - 1] > dp[i][j - 1]`,从`dp[i][j - 1]`结果出发，最多经历1步，即可使`i - 1`与`j - 1`个字符串相等（考虑长度
  `i`字符串经过n步得到`j-1`，`i-1`经过1步能得到`i`，或者反向考虑`j-1`经过n步能得到`i`，最多再经历1步即可由`i`得到`i-1`）
  即`dp[i - 1][j - 1] <= dp[i][j - 1] + 1`

# PrefixSum

* 数组`nums` 构建前缀和数组`preNums`时
    1. 何时构建等长的前缀和数组，即`index`为`n`时，`preNum[n]`代表数组`nums`中当索引`index<=n`时，左侧所有数之和。
    2. 何时构建等长加1的前缀和数组，即`index`为`n`时，`preNum[n]`代表数组`nums`中当索引`index<n`时，左侧所有数之和。
       其实好像都可以，第二种方法更容易处理边界条件，可以比较简便的表示`index=0`时，左侧数之和

# Sliding Window

1. 注意内层for里是满足条件的还是刚刚越过条件的，会有明显的区别

# trick

* big prime:
    * 100000007
    * 1000000007