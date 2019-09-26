import React, { useEffect } from 'react';
import { BudgetContext } from './budget';
import { useQuery, useMutation } from '@apollo/react-hooks';
import PropTypes from 'prop-types';
import { QUERY_BUDGETS, CREATE_BUDGET } from './BudgetProvider.gql';
import { addToList } from '../../../util/immutable';

export const BudgetProvider = ({ children }) => {
  const { loading, error, data } = useQuery(QUERY_BUDGETS);
  const [createBudget] = useMutation(CREATE_BUDGET, {
    update(
      cache,
      {
        data: { createBudget },
      }
    ) {
      const { budgets } = cache.readQuery({ query: QUERY_BUDGETS });
      cache.writeQuery({
        query: QUERY_BUDGETS,
        data: { budgets: addToList(budgets, createBudget) },
      });
    },
  });

  useEffect(() => {
    if (loading || error) return;
    if (!hasAnyBudget(data)) {
      createBudget({ variables: { name: 'default' } });
    }
  }, [loading, error, data, createBudget]);

  if (loading || !hasAnyBudget(data)) return <p>Loading...</p>;
  if (error) {
    console.error(error);
    return <p>Error :(</p>;
  }

  return (
    <BudgetContext.Provider value={data.budgets[0]}>
      {children}
    </BudgetContext.Provider>
  );
};

BudgetProvider.propTypes = {
  children: PropTypes.node,
};

function hasAnyBudget(data) {
  return data && data.budgets && data.budgets.length !== 0;
}
