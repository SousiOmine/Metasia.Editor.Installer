import { useEffect, useState } from 'react';
import '../App.css';
import { StepProps } from '../types/wizard';

function ConfirmationStep({ setCanNavigateNext, setCanNavigateBack, setNextLabel }: StepProps) {
    useEffect(() => {
        setCanNavigateNext(true);
        setCanNavigateBack(true);
        setNextLabel("インストール");
    }, [setCanNavigateNext, setCanNavigateBack, setNextLabel]);
    return (
        <div className="welcome-step">
            <h1>インストールの準備ができました</h1>
        </div>
    )
}

export default ConfirmationStep;