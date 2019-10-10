import React from 'react';
import Page from '../template/Page/Page';
import PageHeader from '../template/Page/PageHeader';
import { EnvelopesListPanel } from './EnvelopesListPanel';
import { CategoriesListPanel } from '../Categories/CategoriesListPanel';

export default function EnvelopesPage() {
  return (
    <Page>
      <PageHeader>Envelopes</PageHeader>
      <EnvelopesListPanel />
      <CategoriesListPanel />
    </Page>
  );
}
