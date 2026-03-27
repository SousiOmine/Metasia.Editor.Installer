import { main } from '../../wailsjs/go/models';

export interface StepProps {
    setCanNavigateNext: (canNavigateNext: boolean) => void;
    setCanNavigateBack: (canNavigateBack: boolean) => void;

    setNextLabel: (label: string) => void;

    context: main.InstallParams;
}

export type StepComponent = React.FC<StepProps>;