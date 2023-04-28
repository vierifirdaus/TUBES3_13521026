export interface message {
    id : number;
    id_histori : number;
    jenis : "input" | "output";
    isi : string;
}

export interface chatProps {
    className: string;
    clickSide : number;
}

export interface messageProps {
    message: message;
}

export interface sidebarProps {
    className: string,
    setClickSide : React.Dispatch<React.SetStateAction<number>>,
    value : string,
    setValue : React.Dispatch<React.SetStateAction<string>>,
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
}

export interface sendMessageProps{
    inputValue: string,
    setInputValue: React.Dispatch<React.SetStateAction<string>>,
    handleInput: (e: React.FormEvent<HTMLFormElement>) => void
}