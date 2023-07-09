module.exports = {
	parserPreset: 'conventional-changelog-conventionalcommits',
	rules: {
		'body-leading-blank': [1, 'always'],
		'body-max-line-length': [2, 'always', 100],
		'footer-leading-blank': [1, 'always'],
		'footer-max-line-length': [2, 'always', 100],
		'header-max-length': [2, 'always', 100],
		'subject-case': [
			2,
			'never',
			['sentence-case', 'start-case', 'pascal-case', 'upper-case'],
		],
		'subject-empty': [2, 'never'],
		'subject-full-stop': [2, 'never', '.'],
		'type-case': [2, 'always', 'lower-case'],
		'type-empty': [2, 'never'],
		'type-enum': [
			2,
			'always',
			[
        'feat',
				'be',
				'fe',
				'devops',
				'db',
				'version',
				'project',
				'openapi',
			],
		],
	},
	prompt: {
		questions: {
			type: {
				description: "Select the type of change that you're committing",
				enum: {
          feat: {
						description: 'A new feature',
						title: 'Feature',
						emoji: '✨',
          },
          be: {
						description: 'A backend related',
						title: 'Backend',
						emoji: '✨',
          },
					fe: {
						description: 'A frontkend related',
						title: 'Frontend',
						emoji: '🐛',
					},
					devops: {
						description: 'A devops related',
						title: 'DevOps',
						emoji: '📚',
					},
					db: {
						description: 'A database related',
						title: 'Database',
						emoji: '💎',
					},
					version: {
						description: 'A version control related',
						title: 'Version Control',
						emoji: '📦',
					},
					project: {
						description: 'Project related',
						title: 'Project',
						emoji: '🚀',
					},
					openapi: {
						description: 'Open Api related',
						title: 'Open API',
						emoji: '🚨',
					},
				},
			},
			scope: {
				description:
					'What is the scope of this change (e.g. component or file name)',
			},
			subject: {
				description:
					'Write a short, imperative tense description of the change',
			},
			body: {
				description: 'Provide a longer description of the change',
			},
			isBreaking: {
				description: 'Are there any breaking changes?',
			},
			breakingBody: {
				description:
					'A BREAKING CHANGE commit requires a body. Please enter a longer description of the commit itself',
			},
			breaking: {
				description: 'Describe the breaking changes',
			},
			isIssueAffected: {
				description: 'Does this change affect any open issues?',
			},
			issuesBody: {
				description:
					'If issues are closed, the commit requires a body. Please enter a longer description of the commit itself',
			},
			issues: {
				description: 'Add issue references (e.g. "fix #123", "re #123".)',
			},
		},
	},
};
