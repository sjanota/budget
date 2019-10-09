import React from 'react';
import Page from '../template/Page/Page';
import PageHeader from '../template/Page/PageHeader';
import { EnvelopesListPanel } from './EnvelopesListPanel';

export default function EnvelopesPage() {
  return (
    <Page>
      <PageHeader>Envelopes</PageHeader>
      <EnvelopesListPanel />
    </Page>
  );
}
