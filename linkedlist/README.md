🟦 单链表题目
1. 反转单链表（Reverse Linked List）
输入：1 → 2 → 3 → 4 → nil

输出：4 → 3 → 2 → 1 → nil


2. 删除指定值节点（Delete Node by Value）
给定一个值 val，删除所有值为 val 的节点。

3. 查找中间节点（Find Middle Node）
若长度为偶数，返回中间偏右的节点。

4. 查找倒数第 k 个节点（Find K-th from End）
输入：1 → 2 → 3 → 4 → 5 → nil, k = 2

输出：4

5. 判断是否有环（Has Cycle）
检测链表中是否有环（使用快慢指针）。

6. 合并两个升序链表（Merge Two Sorted Lists）
输入：两个升序链表

输出：合并后的升序链表

7. 删除倒数第 k 个节点
输入链表和整数 k，删除倒数第 k 个节点。

8. 反转前 N 个节点
输入：1 → 2 → 3 → 4 → 5, N = 3

输出：3 → 2 → 1 → 4 → 5

9. 反转区间 m 到 n 的节点（Reverse Between）
输入：1 → 2 → 3 → 4 → 5, m = 2, n = 4

输出：1 → 4 → 3 → 2 → 5

10. 判断是否为回文链表（Is Palindrome）
不借助数组，使用快慢指针 + 原地反转。

11. 判断两个链表是否相交（Intersection Node）
返回相交的第一个节点（引用相同）。

🟥 双向链表题目（Doubly Linked List）
1. 实现双向链表结构
支持操作：

AddAtHead(val)

AddAtTail(val)

AddAtIndex(index, val)

DeleteAtIndex(index)

PrintForward(), PrintBackward()

2. 设计浏览器历史功能（Browser History）
实现：

Visit(url string)

Back(steps int)

Forward(steps int)

3. 实现 LRU 缓存结构（Least Recently Used Cache）
结合双向链表和哈希表（map）实现

支持：

Get(key)

Put(key, value)

淘汰最久未使用节点

4. 插入排序（Insertion Sort on Doubly Linked List）
使用双向链表实现插入排序

5. 翻转双向链表（Reverse Doubly Linked List）
注意需要同时调整 Prev 和 Next

6. 删除指定节点（Delete Node in Doubly Linked List）
在双向链表中删除给定的节点

🔁 进阶挑战（可选）
- 模拟文本编辑器（Undo/Redo 支持）
利用双向链表实现操作历史和撤销重做功能

- 实现链表版栈 & 队列
Stack：用单链表实现 Push, Pop

Queue：用链表实现 Enqueue, Dequeue

📦 Bonus：链表工具函数（建议封装）
PrintList(head *ListNode)

CreateListFromSlice([]int)

CompareList(a, b *ListNode) bool

Length(head *ListNode)

