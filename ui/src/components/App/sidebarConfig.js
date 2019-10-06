export const sidebarConfig = [
  {
    entries: [{ name: 'Dashboard', to: '/', faIcon: 'fa-tachometer-alt' }],
  },
  {
    name: 'Interface',
    entries: [
      {
        name: 'Components',
        faIcon: 'fa-cog',
        sections: [
          {
            name: 'Custom Components',
            routes: [
              { label: 'Buttons', to: '/buttons' },
              { label: 'Cards', to: '/cards' },
            ],
          },
        ],
      },
      {
        name: 'Utilities',
        faIcon: 'fa-wrench',
        sections: [
          {
            name: 'Custom Utilities',
            routes: [
              { label: 'Colors', to: '/utilities-colors' },
              { label: 'Borders', to: '/utilities-borders' },
              { label: 'Animations', to: '/utilities-animations' },
              { label: 'Other', to: '/utilities-other' },
            ],
          },
        ],
      },
    ],
  },
  {
    name: 'Addons',
    entries: [
      {
        name: 'Pages',
        faIcon: 'fa-folder',
        sections: [
          {
            name: 'Login Screens',
            routes: [
              { label: 'Login', to: '/login' },
              { label: 'Register', to: '/register' },
              { label: 'Forgot Password', to: '/forgot-password' },
            ],
          },
          {
            name: 'Other Pages',
            routes: [
              { label: '404 Page', to: '/404' },
              { label: 'Blank Page', to: '/blank' },
            ],
          },
        ],
      },
      {
        name: 'Charts',
        to: '/charts',
        faIcon: 'fa-chart-area',
      },
      {
        name: 'Tables',
        to: '/tables',
        faIcon: 'fa-table',
      },
    ],
  },
];
