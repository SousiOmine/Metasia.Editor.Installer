import {useState} from 'react';
import logo from './assets/images/logo-universal.png';
import './App.css';
import { Greet } from "../wailsjs/go/main/App";
import WelcomeStep from "./steps/WelcomeStep";
import SelectInstallPathStep from "./steps/SelectInstallPathStep";

function App() {
    const [resultText, setResultText] = useState("あなたの名前を入力してね 👇");
    const [name, setName] = useState('');
    const updateName = (e: any) => setName(e.target.value);
    const updateResultText = (result: string) => setResultText(result);

    const [wizardStep, setWizardStep] = useState(0);

    const renderWizard = () => {
        switch (wizardStep) {
            case 0:
                return <WelcomeStep />
            case 1:
                return <SelectInstallPathStep />
            default:
                return <WelcomeStep />
        }
    }

    function greet() {
        Greet(name).then(updateResultText);
    }

    return (
        <div id="App">
            <div id="result" className="result">{resultText}</div>
            <div id="input" className="input-box">
                <input id="name" className="input" onChange={updateName} autoComplete="off" name="input" type="text"/>
                <button className="btn" onClick={greet}>Greet</button>
            </div>
            {renderWizard()}
            {wizardStep > 0 && <button className="btn" onClick={() => setWizardStep(wizardStep - 1)}>前へ</button>}
            <button className="nextButton" onClick={() => setWizardStep(wizardStep + 1)}>次へ</button>
        </div>
    )
}

export default App
