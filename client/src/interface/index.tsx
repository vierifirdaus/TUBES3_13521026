export interface message {
    id : number;
    id_histori : number;
    jenis : "input" | "output";
    isi : string;
}

export interface chatProps {
    className: string;
    clickSide : number;
    setHistories: React.Dispatch<React.SetStateAction<history[]>>;
    setClicked: React.Dispatch<React.SetStateAction<number>>;
    count: number;
    setCount: React.Dispatch<React.SetStateAction<number>>;
    clicked: number;
    setClickSide: React.Dispatch<React.SetStateAction<number>>;
    chatLog: message[];
    setChatLog: React.Dispatch<React.SetStateAction<message[]>>;
    refHistori:React.RefObject<HTMLDivElement>;

}

export interface messageProps {
    message: message;
}

export interface sidebarProps {
    className: string,
    setClickSide : React.Dispatch<React.SetStateAction<number>>,
    value : string,
    setValue : React.Dispatch<React.SetStateAction<string>>,
    history: history[],
    setHistories: React.Dispatch<React.SetStateAction<history[]>>,
    clicked: number,
    setClicked: React.Dispatch<React.SetStateAction<number>>,
    setCount: React.Dispatch<React.SetStateAction<number>>,
    setChatLog: React.Dispatch<React.SetStateAction<message[]>>,
    refHistori:React.RefObject<HTMLDivElement>;
  }

export interface history{
    id: number,
    nama: string,
}

export interface buttonProps{
    history: history,
    clicked: number,
    handleClick: (i:number) => void
    handleDelete: (i:number) => void
    setHistories: React.Dispatch<React.SetStateAction<history[]>>;
}

export interface sendMessageProps{
    inputValue: string,
    setInputValue: React.Dispatch<React.SetStateAction<string>>,
    handleInput: (e: React.FormEvent<HTMLFormElement> | React.KeyboardEvent<HTMLTextAreaElement>) => void
}