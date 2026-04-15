export function checkOperator(operator: string): boolean {
    return operator in ['+', '-', '/', '*'];
}