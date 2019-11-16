export default {
  buttons: {
    create: 'Dodaj',
    cancel: 'Anuluj',
    save: 'Zapisz',
  },
  sidebar: {
    sections: {
      currentMonth: 'Bieący miesiąc',
      budget: 'Budżet',
    },
    pages: {
      dashboard: 'Podsumowanie',
      expenses: 'Wydatki',
      transfers: 'Przelewy',
      plans: 'Plany',
      accounts: 'Konta',
      envelopes: 'Koperty',
    },
  },
  topbar: {
    budgetLabel: 'Budżet',
    alertsLabel: 'Powiadomienia',
    messagesLabel: 'Wiadomości',
  },
  dashboard: {
    planned: 'Zaplanowano',
    incomes: 'Wpływy',
    leftToPlan: 'Pozostało do zaplanowania',
    expenses: 'Wydatki',
    currentMonth: 'Bieący miesiąc',
    noProblems: 'Wszystko w porządku',
    problems: {
      overplanned: 'Plany na bieżący miesiąc przekraczają wydatki',
      underplanned: 'Część środków jest nie rozplanowana',
      expensesExceedPlans: envelope =>
        `Wydatki przekroczyły zaplanowany budżet na kopercie "${envelope}"`,
      envelopeOverLimit: envelope =>
        `Limit dla koperty "${envelope}" został przekroczony`,
      negativeAccountBalance: account =>
        `Bilans na koncie "${account}" jest ujemny`,
      monthNotEnded: 'Miesiąc się jeszcze nie skończył',
    },
    buttons: {
      closeMonth: 'Zamknij miesiąc',
    },
  },
  accounts: {
    table: {
      title: 'Konta',
      columns: {
        balance: 'Bilans',
        name: 'Nazwa',
      },
    },
    modal: {
      createTitle: 'Dodaj nowe konto',
      editTitle: 'Edytuj konto',
      labels: {
        name: 'Nazwa',
      },
    },
  },
  envelopes: {
    table: {
      title: 'Koperty',
      columns: {
        balance: 'Bilans',
        name: 'Nazwa',
        limit: 'Limit',
        overLimit: 'Ponad limit',
      },
    },
    modal: {
      createTitle: 'Dodaj nową kopertę',
      editTitle: 'Edytuj kopertę',
      labels: {
        name: 'Nazwa',
        limit: 'Limit',
      },
    },
  },
  categories: {
    table: {
      title: 'Kategorie',
      columns: {
        name: 'Nazwa',
        envelope: 'Koperta',
      },
    },
    modal: {
      createTitle: 'Dodaj nową kategorię',
      editTitle: 'Edytuj kategorię',
      labels: {
        name: 'Nazwa',
        envelope: 'Koperta',
      },
    },
  },
  plans: {
    table: {
      title: 'Plany',
      columns: {
        title: 'Tytuł',
        fromEnvelope: 'Z',
        toEnvelope: 'Do',
        currentAmount: 'Kwota',
      },
    },
    modal: {
      createTitle: 'Dodaj nowy plan',
      editTitle: 'Edytuj plan',
      labels: {
        title: 'Tytuł',
        fromEnvelope: 'Z',
        toEnvelope: 'Do',
        amount: 'Kwota',
        recurring: 'Cyklicznie',
      },
    },
  },
  months: [
    'Styczeń',
    'Luty',
    'Marzec',
    'Kwiecień',
    'Maj',
    'Czerwiec',
    'Lipiec',
    'Sierpień',
    'Wrzesień',
    'Październik',
    'Listopad',
    'Grudzień',
  ],
};
