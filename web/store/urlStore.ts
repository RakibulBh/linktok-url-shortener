import { create } from "zustand";

interface iURLStore {
  url: string;
  setUrl: (char: string) => void;
}

interface iResultStore {
  result: string;
  setResult: (char: string) => void;
}

export const useUrlStore = create<iURLStore>((set) => ({
  url: "",
  setUrl: (char) => set((state) => ({ url: char })),
}));

export const useResultStore = create<iResultStore>((set) => ({
  result: "",
  setResult: (char) => set((state) => ({ result: char })),
}));
