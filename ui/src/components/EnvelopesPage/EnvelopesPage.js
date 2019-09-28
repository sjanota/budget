import React from 'react';
import { EnvelopesList } from '../EnvelopesList/EnvelopesList';
import { CategoriesList } from '../CategoriesList/CategoriesList';
import './EnvelopesPage.css';

export function EnvelopesPage() {
  return (
    <div className="EnvelopesPage">
      <EnvelopesList />
      <CategoriesList />
    </div>
  );
}
