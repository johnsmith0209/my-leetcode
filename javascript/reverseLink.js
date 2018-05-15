// TO BE FINISHED
/**
 * Definition for singly-linked list.
 * function ListNode(val) {
 *     this.val = val;
 *     this.next = null;
 * }
 */

function ListNode(val, next = null) {
  this.val = val;
  this.next = next;
}

var dumpList = (head) => {
  let a = []
  while(head) {a.push(head.val); head = head.next}
  console.log(a.join('->'))
}

/**
 * https://leetcode-cn.com/explore/interview/card/top-interview-questions-easy/6/linked-list/43/
 * @param {ListNode} head
 * @return {ListNode}
 */
var reverseList = function(head) {
  if (!head || !head.next) return head
  let current = null;
  let next = head;
  while (next) {
    let tempNext = next.next;
    next.next = current;
    current = next;
    next = tempNext;
  }
  return next;
};

const node1 = new ListNode(5)
const node2 = new ListNode(2, node1)
const node3 = new ListNode(1, node2)
const node4 = new ListNode(6, node3)
dumpList(node4)
reverseList(node4)
dumpList(node1)