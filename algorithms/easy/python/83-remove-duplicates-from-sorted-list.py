# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next
class Solution:
    def deleteDuplicates(self, head: Optional[ListNode]) -> Optional[ListNode]:
        current_node = head
        next_node = current_node
        while next_node:
            if current_node.val != next_node.val:
                current_node.next = next_node
                current_node = current_node.next
            next_node = next_node.next
        if current_node:
            current_node.next = next_node
        return head

        