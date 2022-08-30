# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next
class Solution:
    def reverseBetween(self, head: Optional[ListNode], left: int, right: int) -> Optional[ListNode]:
        stack = []
        i = 1
        node = head
        while i < left:
            node = node.next
            i += 1
        left_node = node
        while i <= right:
            stack.append(node.val)
            node = node.next
            i += 1
        node = left_node
        i = left
        while i <= right:
            node.val = stack.pop()
            node = node.next
            i += 1
        return head
