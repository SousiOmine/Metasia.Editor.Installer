import { useState, useEffect } from 'react';
import { main } from '../wailsjs/go/models';
import './App.css';
import WelcomeStep from "./steps/WelcomeStep";
import SelectInstallPathStep from "./steps/SelectInstallPathStep";
import SelectPluginsStep from "./steps/SelectPluginsStep";
import ConfirmationStep from "./steps/ConfirmationStep";
import InstallationStep from "./steps/InstallationStep";
import { StepProps } from './types/wizard';
import { GetDefaultInstallParams, Quit } from '../wailsjs/go/main/App';

function App() {
    const [resultText, setResultText] = useState("あなたの名前を入力してね 👇");
    const [wizardStep, setWizardStep] = useState(0);
    const [canNavigateNext, setCanNavigateNext] = useState(true);
    const [canNavigateBack, setCanNavigateBack] = useState(true);
    const [nextLabel, setNextLabel] = useState("次へ");

    const [context, setContext] = useState<main.InstallParams>(new main.InstallParams());
    const [initError, setInitError] = useState<string | null>(null);

    useEffect(() => {
        let isMounted = true;

        GetDefaultInstallParams()
            .then((params) => {
                if (isMounted) {
                    setContext(params);
                }
            })
            .catch((err) => {
                if (isMounted) {
                    setInitError(err?.toString() || '初期化に失敗しました');
                    console.error('GetDefaultInstallParams failed:', err);
                }
            });

        return () => {
            isMounted = false;
        };
    }, []);

    const stepProps: StepProps = {
        setCanNavigateNext: setCanNavigateNext,
        setCanNavigateBack: setCanNavigateBack,
        setNextLabel: setNextLabel,
        context: context,
    }

    const steps = [
        <WelcomeStep {...stepProps} />,
        <SelectInstallPathStep {...stepProps} />,
        <SelectPluginsStep {...stepProps} />,
        <ConfirmationStep {...stepProps} />,
        <InstallationStep {...stepProps} />,
    ]

    if (wizardStep >= steps.length) {
        Quit();
    }

    if (initError) {
        return (
            <div id="App">
                <div className="error-container">
                    <h2>初期化エラー</h2>
                    <p>{initError}</p>
                </div>
            </div>
        );
    }

    return (
        <div id="App">
            {steps[wizardStep]}
            <div className='wizard-navigate-button-container'>
                {
                    <button onClick={() => setWizardStep(wizardStep - 1)} aria-disabled={!canNavigateBack}>
                        戻る
                    </button>
                }
                {
                    <button onClick={() => setWizardStep(wizardStep + 1)} aria-disabled={!canNavigateNext}>
                        {nextLabel}
                    </button>
                }
            </div>
        </div>
    )
}

export default App
