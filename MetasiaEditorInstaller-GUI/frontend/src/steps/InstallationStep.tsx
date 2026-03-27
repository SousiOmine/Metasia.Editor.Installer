import { useState, useEffect, useRef } from 'react';
import '../App.css';
import { StepProps } from '../types/wizard';
import { InstallExecute } from '../../wailsjs/go/main/App';


function InstallationStep({ setCanNavigateNext, setCanNavigateBack, setNextLabel, context }: StepProps) {
    const [status, setStatus] = useState("インストール中...");
    const [errorDetail, setErrorDetail] = useState<string | null>(null);
    const hasExecuted = useRef(false);

    useEffect(() => {
        if (hasExecuted.current) return;
        hasExecuted.current = true;
        setCanNavigateNext(false);
        setCanNavigateBack(false);
        setNextLabel("終了");

        const handleInstall = async () => {
            try {
                await InstallExecute(context);
                setStatus("インストールが完了しました。");
                setCanNavigateNext(true);
            } catch (error) {
                setStatus("インストールに失敗しました。");
                setErrorDetail(String(error));
                console.error("Installation failed:", error);
            }
        };

        handleInstall();
    }, [setCanNavigateNext, setCanNavigateBack, setNextLabel, context]);

    return (
        <div className="welcome-step">
            <h1>{status}</h1>
            {errorDetail && (
                <div className='error'>
                    <p>エラー詳細:</p>
                    <pre>{errorDetail}</pre>
                </div>
            )}
        </div>
    )
}

export default InstallationStep;