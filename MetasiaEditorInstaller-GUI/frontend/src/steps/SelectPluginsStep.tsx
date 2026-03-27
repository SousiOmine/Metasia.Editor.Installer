import { useEffect, useState } from 'react';
import '../App.css';
import { StepProps } from '../types/wizard';

function SelectPluginsStep({ setCanNavigateNext, setCanNavigateBack, setNextLabel }: StepProps) {
    useEffect(() => {
        setCanNavigateNext(true);
        setCanNavigateBack(true);
        setNextLabel("次へ");
    }, [setCanNavigateNext, setCanNavigateBack, setNextLabel]);
    return (
        <div className="welcome-step">
            <h1>インストールするプラグインを選択してください</h1>
            まだ選択できません
        </div>
    )
}

export default SelectPluginsStep;