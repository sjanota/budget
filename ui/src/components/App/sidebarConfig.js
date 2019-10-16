export const sidebarConfig = [
  {
    entries: [
      {
        name: 'Dashboard',
        faIcon: 'receipt',
        to: '/',
      },
    ],
  },
  {
    name: 'Current month',
    entries: [
      {
        name: 'Expenses',
        faIcon: 'receipt',
        to: '/expenses',
      },
      {
        name: 'Transfers',
        faIcon: 'exchange-alt',
        to: '/transfers',
      },
      {
        name: 'Plans',
        faIcon: 'map-marked-alt',
        to: '/plans',
      },
    ],
  },
  {
    name: 'Budget',
    entries: [
      {
        name: 'Accounts',
        faIcon: 'credit-card',
        to: '/accounts',
      },
      {
        name: 'Envelopes',
        faIcon: 'envelope-open-text',
        to: '/envelopes',
      },
    ],
  },
];
