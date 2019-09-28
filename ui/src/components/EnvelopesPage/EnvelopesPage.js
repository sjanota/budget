import React from 'react';
import { EnvelopesList } from '../EnvelopesList/EnvelopesList';
import { CategoriesList } from '../CategoriesList/CategoriesList';

export function EnvelopesPage() {
  return (
    <div className="EnvelopesPage">
      <EnvelopesList />
      <CategoriesList />
    </div>
  );
}
