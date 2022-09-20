# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next
class Solution:
    def mergeTwoLists(self, list1: Optional[ListNode], list2: Optional[ListNode]) -> Optional[ListNode]:
        if list1 or list2:
            start_node = ListNode()
            current_node = start_node
            while list1 or list2:
                if not (list1 and list2):
                    if list1:
                        current_node.val = list1.val
                        list1 = list1.next
                    if list2:
                        current_node.val = list2.val
                        list2 = list2.next
                else:
                    if list1.val <= list2.val:
                        current_node.val = list1.val
                        list1 = list1.next
                    else:
                        current_node.val = list2.val
                        list2 = list2.next
                if list1 or list2:
                    current_node.next = ListNode()
                    current_node = current_node.next
            return start_node