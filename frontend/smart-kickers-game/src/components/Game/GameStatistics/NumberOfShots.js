import React from 'react';
import { TeamID } from '../../../constants/score';
import './GameStatistics.css';

function NumberOfShots({ statistics }) {
  return (
    <>
      <div className="table-item">{statistics ? statistics?.teamID[TeamID.Team_blue]?.shotsCount : '0'}</div>
      <div className="table-item">number of all shots in the game</div>
      <div className="table-item">{statistics ? statistics?.teamID[TeamID.Team_white]?.shotsCount : '0'}</div>
    </>
  );
}

export default NumberOfShots;
