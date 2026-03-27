import { useEffect, useState } from 'react';
import '../App.css';
import { StepProps } from '../types/wizard';
import { SelectDirectoryDialog } from '../../wailsjs/go/main/App';

function SelectInstallPathStep({ setCanNavigateNext, setCanNavigateBack, setNextLabel, context }: StepProps) {
    const [selectedPath, setSelectedPath] = useState(context.Path || "");

    useEffect(() => {
        setCanNavigateNext(false);
        setCanNavigateBack(true);
        setNextLabel("次へ");
    }, [setCanNavigateNext, setCanNavigateBack, setNextLabel]);

    const handleSelectPath = async () => {
        try
        {
            const path = await SelectDirectoryDialog(selectedPath);
            if (path) {
                setSelectedPath(path);
                setCanNavigateNext(true);
                context.Path = path;
                context.PluginsPath = path + "/Plugins";
            }
        }
        catch (error)
        {
            console.error("Failed to select directory:", error);
        }
    }

    if (context.Path != "")
    {
        setCanNavigateNext(true);
    }

    return (
        <div className="welcome-step">
            <h1>インストール先を指定してください</h1>
            <div className='path-selector'>
                <input id='install-path' type="text" value={selectedPath} readOnly
                    placeholder='フォルダを選択...'/>
                <button id='select-path-button' onClick={handleSelectPath}>参照</button>
            </div>
        </div>
    )
}


export default SelectInstallPathStep;