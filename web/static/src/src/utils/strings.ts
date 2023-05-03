export function padLeft(num: number) {
  return num.toString().padStart(2, '0')
}

export function newlines(str: string): string {
  return str.replace(/\n/g, '<br>')
}
