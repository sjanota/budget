import React, {useEffect} from 'react';
import {useQuery} from "@apollo/react-hooks";
import {EXPENSES_QUERY, EXPENSES_EVENTS_SUBSCRIPTION} from "./App.gql";

function handleExpenseEvent(prev, { subscriptionData }) {
  const event = subscriptionData.data.expenses;
  switch (event.type) {
    case "ADDED": {
      const newExpense = event.expense;
      return {expenses: [...prev.expenses, newExpense]};
    }
    default:
      return prev;
  }
}

export default function App() {
  const {loading, error, data, subscribeToMore} = useQuery(EXPENSES_QUERY);

  useEffect(() => {
    if (loading) return;
    return subscribeToMore({
      document: EXPENSES_EVENTS_SUBSCRIPTION,
      updateQuery: handleExpenseEvent,
      onError: console.error
    });
  }, [loading, subscribeToMore]);

  if (loading) return <p>Loading...</p>;
  if (error) {
    console.error(error);
    return <p>Error :(</p>;
  }

  return (
    <div className="App">
      {data.expenses.map(e => <pre key={e.id}>{JSON.stringify(e)}</pre>)}
    </div>
  );
};