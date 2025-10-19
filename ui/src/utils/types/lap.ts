export interface Lap {
  // Header
  SessionUID: number;
  PlayerCarIndex: number;

  // // Lap data
  CurrentLapNum: number;

  Sector1Minutes: number; // Sector 1 whole minute part
  Sector1MS: number; // Sector 1 time milliseconds part

  Sector2Minutes: number; // Sector 2 whole minute part
  Sector2MS: number; // Sector 2 time milliseconds part

  Sector3Minutes: number; // Sector 3 whole minute part
  Sector3MS: number; // Sector 3 time milliseconds part

  Total: number;

  CurrentLapInvalid: number; // Current lap invalid - 0 = valid, 1 = invalid

  // Session
  SessionType: number;
  TrackId: number;
}
