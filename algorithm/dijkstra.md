# dijkstra算法

[https://houbb.github.io/2020/01/23/data-struct-learn-03-graph-dijkstra](https://houbb.github.io/2020/01/23/data-struct-learn-03-graph-dijkstra)
> 核心在两点
> * 最短路径的子路径仍然是最短路径(从已经确定是最短的子路径`(D,E...)`，确定从起点经过该点`(D,E...)`到目标点的最短距离，即为目的最短路径)，dijkstra算法的实现从已确定的两点间最短距离的临时终点`(D,E...)`，计算经该点`(D,E...)`到其他可达节点的距离，即为上述逻辑的体现
> * 如何确定两点间的最短路径，在没有负权边的情况下，从起点到该未处理的节点`X`的路径长度最短即可认为是两点间的最短距离`L`（之前已处理节点，及距离小于`L`的节点，不能经其到达`X`以获得更短的路径；距离大于`L`的未处理节点，由于”无负权边“的限制，保证了不存在经其到达X获得更短路径长度的可能）
> * 所以，提供一种思路，为什么这么做的原因，1. 就是找起点到某一个点最短路径；2. 为什么需要1，因为<font color=LightCoral>”最短路径的子路径仍然是最短路径“</font>，若“最短路径”，则“其子路径仍然是最短路径”（前者是后者的充分不必要条件）。同时也包括了这样一个信息，我们需要获得的这条子路径一定在所有<font color=LightCoral>“子路径为最短路径”</font>的范围之中，算法做的就是从所有<font color=LightCoral>“子路径为最短路径”</font>的可能中找到目的路径。<br>

由于这两点减少了遍历的复杂度， 有点动态规划加贪心的感觉

> 为什么dijkstra算法不适用于负权边
> * 原因是在确定最短路径时，<font color=LightCoral>“之前已处理节点，及距离小于`L`的节点，不能经其到达`X`以获得更短的路径；距离大于`L`的未处理节点，由于”无负权边“的限制，保证了不存在经其到达X获得更短路径长度的可能”</font>这个思想。若大于`L`的未处理节点存在负权边，可能会得到更短路径，然而实际处理时会分为两个集合，默认是最短路径不再考虑，所以不能处理。
> * 最直观的考虑怎么解决这个问题，这个最短路径不一定为真，若存在负权边重新遍历
> * Floyd 算法