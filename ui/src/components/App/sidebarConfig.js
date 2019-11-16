export const sidebarConfig = ({ sidebar }) => [
  {
    entries: [
      {
        name: sidebar.pages.dashboard,
        faIcon: 'receipt',
        to: '/',
      },
    ],
  },
  {
    name: sidebar.sections.currentMonth,
    entries: [
      {
        name: sidebar.pages.expenses,
        faIcon: 'receipt',
        to: '/expenses',
      },
      {
        name: sidebar.pages.transfers,
        faIcon: 'exchange-alt',
        to: '/transfers',
      },
      {
        name: sidebar.pages.plans,
        faIcon: 'map-marked-alt',
        to: '/plans',
      },
    ],
  },
  {
    name: sidebar.sections.budget,
    entries: [
      {
        name: sidebar.pages.accounts,
        faIcon: 'credit-card',
        to: '/accounts',
      },
      {
        name: sidebar.pages.envelopes,
        faIcon: 'envelope-open-text',
        to: '/envelopes',
      },
    ],
  },
];
