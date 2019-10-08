export const sidebarConfig = [
  {
    entries: [{ name: 'Dashboard', to: '/', faIcon: 'tachometer-alt' }],
  },
  {
    name: 'Budget',
    entries: [
      {
        name: 'Accounts',
        faIcon: 'credit-card',
        to: '/accounts',
      },
    ],
  },
  {
    name: 'Interface',
    entries: [
      {
        name: 'Components',
        faIcon: 'cog',
        sections: [
          {
            name: 'Custom Components',
            routes: [{ label: 'Buttons', to: '/buttons' }],
          },
        ],
      },
      {
        name: 'Tables',
        to: '/tables',
        faIcon: 'table',
      },
    ],
  },
];
