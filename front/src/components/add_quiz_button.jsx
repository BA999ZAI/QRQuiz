import React from "react";

const AddQuizButton = ({ onClick }) => {
  return (
    <button className="button-auth auth mr-10 cursor-pointer" onClick={onClick}>
      Создать
    </button>
  );
};

export default AddQuizButton;