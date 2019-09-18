import React from 'react';
import {useQuery} from "@apollo/react-hooks";
import {gql} from 'apollo-boost';


const EXPENSE_FRAGMENT = gql`
    fragment ExpensesDetails on Expense {
        id
        title
        date
        total
        location
        account {
            id
            name
            available
        }
        entries {
            title
            amount
            category {
                id
                name
            }
        }
    }
`;

const QUERY_EXPENSES = gql`
    query QueryExpenses {
        expenses {
            ...ExpensesDetails
        }
    }
    ${EXPENSE_FRAGMENT}
`;

const SUBSCRIBE_EXPENSES = gql`
    subscription WatchExpenses {
        expenses {
            type
            ... on ExpenseAdded {
                expense {
                    ...ExpensesDetails
                }
            }
        }
    }
    ${EXPENSE_FRAGMENT}
`;

function App() {
  const {loading, error, data, subscribeToMore} = useQuery(QUERY_EXPENSES);

  React.useEffect(() => {
    return subscribeToMore({
      document: SUBSCRIBE_EXPENSES,
      updateQuery: (prev, { subscriptionData }) => {
        if (!subscriptionData.data) return prev;
        switch (subscriptionData.data.expenses.type) {
          case "ADDED": {
            const newExpense = subscriptionData.data.expenses.expense;
            return {expenses: [...prev.expenses, newExpense]};
          }
          default:
            return prev;
        }
      },
      onError: console.error
    });
  }, [subscribeToMore]);


  if (loading) return <p>Loading...</p>;
  if (error) {
    console.warn(error);
    return <p>Error :(</p>;
  }

  console.log(loading, error, data);

  return (
    <div className="App">
      <pre>{JSON.stringify(data, null, 2)}</pre>
    </div>
  );
}

export default App;
