import gql from 'graphql-tag';
import { useMutation, useQuery } from '@apollo/react-hooks';
import { useBudget } from './budget';

const GET_CATEGORIES = gql`
  query GetCategories($budgetID: ID!) {
    categories(budgetID: $budgetID) {
      id
      name
      envelope {
        id
        name
      }
    }
  }
`;

const CREATE_CATEGORY = gql`
  mutation CreateCategory($budgetID: ID!, $input: CategoryInput!) {
    createCategory(budgetID: $budgetID, in: $input) {
      id
      name
      envelope {
        id
        name
      }
    }
  }
`;

const UPDATE_CATEGORY = gql`
  mutation UpdateCategory($budgetID: ID!, $id: ID!, $input: CategoryUpdate!) {
    updateCategory(budgetID: $budgetID, id: $id, in: $input) {
      id
      name
      envelope {
        id
        name
      }
    }
  }
`;

export function useCreateCategory() {
  const { selectedBudget } = useBudget();
  const [mutation, ...rest] = useMutation(CREATE_CATEGORY, {
    update: (cache, { data: { createCategory } }) => {
      const { categories } = cache.readQuery({
        query: GET_CATEGORIES,
        variables: { budgetID: selectedBudget.id },
      });
      cache.writeQuery({
        query: GET_CATEGORIES,
        variables: { budgetID: selectedBudget.id },
        data: {
          categories: categories.concat([createCategory]),
        },
      });
    },
  });
  const wrapper = input => {
    mutation({ variables: { budgetID: selectedBudget.id, input } });
  };
  return [wrapper, ...rest];
}

export function useUpdateCategory() {
  const { selectedBudget } = useBudget();
  const [mutation, ...rest] = useMutation(UPDATE_CATEGORY);
  const wrapper = (id, input) => {
    mutation({ variables: { budgetID: selectedBudget.id, id, input } });
  };
  return [wrapper, ...rest];
}

export function useGetCategories() {
  const { selectedBudget } = useBudget();
  return useQuery(GET_CATEGORIES, {
    variables: { budgetID: selectedBudget.id },
  });
}
