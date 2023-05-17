/*
 * For a detailed explanation regarding each configuration property and type check, visit:
 * https://jestjs.io/docs/configuration
 */

export default {
  // Automatically clear mock calls, instances, contexts and results before every test
  clearMocks: true,

  // Indicates whether the coverage information should be collected while executing the test
  collectCoverage: true,

  // An array of glob patterns indicating a set of files for which coverage information should be collected
  // collectCoverageFrom: undefined,

  // The directory where Jest should output its coverage files
  coverageDirectory: "coverage",
  transform: {
    "^.+\\.(ts|js)x?$": [
      "ts-jest",
      {
        diagnostics: false,
      },
    ],
  },
  moduleNameMapper: {
    "~/(.*)": "<rootDir>/app/$1",
  },
};
