import { useEffect, useState } from 'react';
import './App.css';
import { resetGame } from './apis/Game';
import { Button } from './components/Button';
import GameResults from './components/GameResults.js';
import Heatmap from './components/Heatmap/Heatmap';

import config from './config';
import { getHeatmapData } from './apis/heatmap';

function App() {
  const [blueScore, setBlueScore] = useState(0);
  const [whiteScore, setWhiteScore] = useState(0);
  const [clicked, setClicked] = useState(0);
  const [heatmap, setHeatmap] = useState([]);

  useEffect(() => {

    const socket = new WebSocket(`${config.wsBaseUrl}/score`);

    socket.onopen = function () {
      // Send to server
      socket.send('Hello from client');
      socket.onmessage = (msg) => {
        msg = JSON.parse(msg.data);
        setBlueScore(msg.blueScore);
        setWhiteScore(msg.whiteScore);
      };
    };
  }, []);

  function handleResetGame() {
    resetGame().then((data) => {
      if (data.error) alert(data.error);
    });
  }

  let heatMapTable = []
  async function getHeatmap() {
    heatMapTable = await getHeatmapData()
    setClicked(true)
    console.log(heatMapTable)
  }

  

  return (
    <>
      <h1>Smart Kickers</h1>
      <GameResults blueScore={blueScore} whiteScore={whiteScore} />
      <center>
        <Button onClick={() => handleResetGame()}>Reset game</Button>
        <Button onClick={() => getHeatmap()}>Stats</Button>
      </center>
      {clicked ? (<Heatmap heatMapTable={heatMapTable}/>) : null}
      {/* <Heatmap heatMapTable={heatMapTable}/> */}
    </>
  );
}

export default App;
