export interface Config {
  Port: number;
  TelemetryFileName: string;
  LapsFileName: string;
  UseObs: number;
  ObsBufferSeconds: number;
  ObsPassword: string;
  ObsPort: number;
  ObsAddr: string;
  TelemetryFMBufferRows: number;
}
