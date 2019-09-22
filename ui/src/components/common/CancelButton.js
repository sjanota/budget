import React from 'react';
import { X } from "@primer/octicons-react";
import { OcticonButton } from './OcticonButton';
import PropTypes from 'prop-types'

export function CancelButton({ onClick }) {
    return <OcticonButton icon={X} action={"cancel"} onClick={onClick} />;
}

CancelButton.propTypes = {
    onClick: PropTypes.func
}
