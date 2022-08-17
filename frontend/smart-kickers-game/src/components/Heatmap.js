import * as React from 'react';
import CalendarHeatmap from 'react-calendar-heatmap';
import 'react-calendar-heatmap/dist/styles.css';
import './Heatmap.css';
const Heatmap = () => {
  return (
    <CalendarHeatmap
      showMonthLabels={false}
      showWeekdayLabels={false}
      values={[
        { date: shiftDate(2, -1), count: 1 },
        { count: 1 },
        { count: 11 },
        { count: 11 },
        { count: 11 },
        { count: 1114 },
        { count: 1114 },
        { count: 2 },
        // ...and so on
      ]}
      classForValue={(value) => {
        if (!value) {
          return 'color-empty';
        }
        return `color-scale-${value.count}`;
      }}
    ></CalendarHeatmap>
  );
};
function shiftDate(date, numDays) {
  const newDate = new Date(date);
  newDate.setDate(newDate.getDate() + numDays);
  return newDate;
}

export default Heatmap;
