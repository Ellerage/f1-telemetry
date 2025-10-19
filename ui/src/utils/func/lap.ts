export function getSectorString(minutes: number, ms: number): string {
  return `${minutes}:${Math.floor(ms / 1000)}.${ms % 1000}`;
}

export function getLapTime(ms: number): string {
  return `${Math.floor(ms / 60_000)}:${Math.floor((ms % 60_000) / 1000)}.${
    ms % 1000
  }`;
}
