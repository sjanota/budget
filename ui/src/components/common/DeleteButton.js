import React from 'react';
import { Trashcan } from "@primer/octicons-react";
import PropTypes from 'prop-types';
import { OcticonButton } from './OcticonButton';

export function DeleteButton({ onClick }) {
    return <OcticonButton icon={Trashcan} action={"delete"} onClick={onClick} />;
}

DeleteButton.propTypes = {
    onClick: PropTypes.func
};
