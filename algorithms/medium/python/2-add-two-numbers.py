# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next
class Solution:
    def addTwoNumbers(self, l1: Optional[ListNode], l2: Optional[ListNode]) -> Optional[ListNode]:
        start_node = ListNode()
        current_node = start_node
        while l1 or l2:
            if l1:
                current_node.val += l1.val
                l1 = l1.next
            if l2:
                current_node.val += l2.val
                l2 = l2.next
            if current_node.val > 9:
                current_node.next = ListNode()
                current_node.next.val = 1
                current_node.val %= 10
            elif l1 or l2:
                current_node.next = ListNode()
            current_node = current_node.next
        current_node = None
        return start_node