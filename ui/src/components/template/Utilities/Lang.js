import React, { createContext, useContext } from 'react';

const LangContext = createContext();

export const LangProvider = ({ dictionary, children }) => (
  <LangContext.Provider value={dictionary}>{children}</LangContext.Provider>
);

export const useDictionary = () => useContext(LangContext);
