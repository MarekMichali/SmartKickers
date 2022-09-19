import React from 'react';
import './StatisticItem.css';

function ManualChangedGoals({ blue, white, statistics }) {
  function getManualSubstractedGoals(teamID) {
    return statistics?.teamID[teamID]?.manualGoals['sub'] || 0;
  }

  function getManualAddedGoals(teamID) {
    return statistics?.teamID[teamID]?.manualGoals['add'] || 0;
  }
  return (
    <>
      <div className="table-item">{getManualAddedGoals(blue)}</div>
      <div className="table-item">Manually added goals</div>
      <div className="table-item">{getManualAddedGoals(white)}</div>
      <div className="table-item">{getManualSubstractedGoals(blue)}</div>
      <div className="table-item">Manually substracted goals</div>
      <div className="table-item">{getManualSubstractedGoals(white)}</div>
    </>
  );
}

export default ManualChangedGoals;
