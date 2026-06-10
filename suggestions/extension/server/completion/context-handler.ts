import * as fs from 'fs';
import * as path from 'path';
import { handleApply } from "../rules/apply";
import { handleIdentifier } from "../rules/identifier";
import { handleOperation } from "../rules/normal-operation";
import { handleWhenOperation } from "../rules/when-operation";

const logPath = path.join(__dirname, 'PROVIDER_LOG.txt');

function logToFile(message: any) {
  const text = typeof message === 'string' ? message : JSON.stringify(message, null, 2);
  fs.appendFileSync(logPath, text + '\n');
}

export function handleNext(children: any[], cursorOffset: number) {
  let leftNode = null;
  const types = Array.from(children, (child: any) => child.type);
  const firstType = types[0];

  for (const child of children) {
    if (child.endIndex <= cursorOffset) {
      if (!leftNode || child.endIndex > leftNode.endIndex) {
        leftNode = child;
      }
    }
  }

  if (!leftNode) {
    return [];
  }

  const leftNodeIndex = children.indexOf(leftNode);
  const prevNode = leftNodeIndex > 0 ? children[leftNodeIndex - 1] : null;

  logToFile(`Left: ${leftNode.type} | Prev: ${prevNode?.type}`);

  switch (firstType) {
    case "APPLY":
    case "apply":
      return handleApply(prevNode, leftNode);

    case "identifier":
      return handleIdentifier(prevNode, leftNode);

    case "EXPLAIN":
    case "explain":
      return handleOperation(prevNode, leftNode);

    case "WHEN":
    case "when":
      return handleWhenOperation(prevNode, leftNode);

    case "=":
    case "=>":
      return [];
  }

  return [];
}