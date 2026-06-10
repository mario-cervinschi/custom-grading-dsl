export function findRelevantNodes(tree: any, offset: number) {
  let node = tree.rootNode.descendantForIndex(Math.max(0, offset));

  if (node.type === "source_file" || node.type === "program") {
    let bestChild = null;
    let bestEnd = -1;
    for (let i = 0; i < node.childCount; i++) {
      const child = node.child(i);
      if (child && child.endIndex <= offset) {
        if (child.endIndex > bestEnd) {
          bestChild = child;
          bestEnd = child.endIndex;
        }
      }
    }

    if (bestChild) node = bestChild;
  }

  while (node && node.type === "ERROR") {
    if (node.childCount > 0) {
      node = node.lastChild;
    } else {
      break;
    }
  }

  const candidates = [];
  const queue = [];

  if (node) {
    candidates.push(node);
    queue.push(node);

    let notDone = true;
    while (notDone) {
      const currentNode: any = queue.pop();

      if (!currentNode) {
        notDone = false;
        break;
      }

      for (let i = 0; i < currentNode.childCount; i++) {
        candidates.push(currentNode.child(i));
        queue.push(currentNode.child(i));
      }

      if (queue.length === 0) {
        notDone = false;
        break;
      }
    }
  }

  return { node, candidates };
}
