box.cfg {
    listen = 3301
}

require('console').start()

o = box.schema.space.create('organisation', {if_not_exists = true})
o:format({
{name = 'id', type = 'unsigned'},
{name = 'name', type = 'string'},
{name = 'status', type = 'boolean'},
})

box.schema.sequence.create('orgS',{if_not_exists = true})

o:create_index('primary', {if_not_exists = true,
sequence='orgS',
type = 'tree',
parts = {'id'}
})

u = box.schema.space.create('user', {if_not_exists = true})
u:format({
{name = 'id', type = 'unsigned'},
{name = 'name', type = 'string'},
{name = 'org_id', type = 'unsigned'}
})

box.schema.sequence.create('userS',{if_not_exists = true})

u:create_index('primary', {
if_not_exists = true,
sequence='userS',
type = 'tree',
parts = {'id'}
})

f = box.schema.space.create('file', {if_not_exists = true})
f:format({
{name = 'id', type = 'uuid'},
{name = 'name', type = 'string'},
{name = 'location', type = 'string'},
{name = 'user_id', type = 'unsigned'},
{name = 'external_uuid', type = 'uuid'}
})

f:create_index("pk", {if_not_exists = true, parts={{field = 1, type = 'uuid'}}})
