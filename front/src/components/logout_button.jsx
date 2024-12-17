import React from "react";

const LogoutButton = ({ onClick }) => {
  return (
    <a className="button-auth auth mr-10 cursor-pointer" onClick={onClick}>
      Выйти
    </a>
  );
};

export default LogoutButton;