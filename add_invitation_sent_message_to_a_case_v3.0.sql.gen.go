// AUTOMATICALLY GENERATED FILE. DO NOT EDIT.

package invite

var add_invitation_sent_message_to_a_case_v3 = esql(asset{Name: "add_invitation_sent_message_to_a_case_v3.0.sql", Content: "" +
	"# For any question about this script, ask Franck\n#\n#################################################################\n#\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t#\n# UPDATE THE BELOW VARIABLES ACCORDING TO YOUR NEEDS\t\t\t#\n#\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t#\n#################################################################\n#\n# The MEFE invitation id that we want to finalize:\n\tSET @mefe_invitation_id = '%s';\n\tSET @mefe_invitation_id_int_value = %d;\n\tSET @environment = %d;\n#\n# The Timestamp when the MEFE invitation was sent:\n\tSET @mefe_invitation_sent = NOW();\n#\n########################################################################\n#\n#\tALL THE VARIABLES WE NEED HAVE BEEN DEFINED, WE CAN RUN THE SCRIPT #\n#\n########################################################################\n#\n#############################################\n#\t\t\t\t\t\t\t\t\t\t\t#\n# IMPORTANT INFORMATION ABOUT THIS SCRIPT\t#\n#\t\t\t\t\t\t\t\t\t\t\t#\n#############################################\n#\n# Built for BZFE database v3.0\n#\n# This assumes that the invitation has been:\n#\t- Created in the MEFE\n#\t- Processed in the BZ DB\n#\t- Finalized in the MEFE (an email has been sent to the invited user)\n#\n# Pre-requisite:\n#\t- The table 'ut_invitation_api_data' has been updated\n# \t- We know the MEFE Invitation id that we need to process.\n#\t- We know the environment where this script is run\n# \n# This script depends on one SQL procedure 'finalize_invitation_to_a_case'.\n# This procedure needs the following variables:\n#\t- @mefe_invitation_id\n#\t- @mefe_invitation_sent\n#\t- @bz_case_id\n#\t- @bz_user_id\n#\t- @creator_bz_id\n#\t- @user_role_type_name\n#\t- @mefe_invitor_user_id\n#\n# This script assumes that:\n#\t- An invitation has been created in the MEFE\n#\t- The invitation has been processed in the BZ database\n#\t- The MEFE has been notified that the invitation has been processed in the BZ database and has sent an email to the invitee.\n#\n# This script will:\n#\t- Populate the needed variables\n#\t- Verify that the invitation has been correctly processed in the BZ database\n# \t- Call the procedure to finalize the invitation\n#\t\t- Add a comment to the case to indicate that the invitation has been sent\n#\t\t- Log its actions\n#\t\t- Update the table 'ut_data_to_add_user_to_a_case' for future reference and audit.\n#\n#\n#####################################################\n#\t\t\t\t\t\n# First we need to define all the variables we need\n#\t\t\t\t\t\n#####################################################\n\n# Info about this script\n\tSET @this_script = 'add_invitation_sent_message_to_a_case_v3.0.sql';\n\n# Timestamp\t\n\tSET @timestamp = NOW();\n\t\n# The reference of the record we want to update in the table ''\n\tSET @reference_for_update = (SELECT `id` FROM `ut_invitation_api_data` WHERE `mefe_invitation_id` = @mefe_invitation_id);\t\n\n# The case id\n\tSET @bz_case_id = (SELECT `bz_case_id` FROM `ut_invitation_api_data` WHERE `id` = @reference_for_update);\n\t\n# The user who you want to associate to this unit - BZ user id of the user that you want to associate/invite to the unit.\n\tSET @bz_user_id = (SELECT `bz_user_id` FROM `ut_invitation_api_data` WHERE `id` = @reference_for_update);\n\n# The Invitor - BZ user id of the user that has genereated the invitation.\n\tSET @creator_bz_id = (SELECT `bzfe_invitor_user_id` FROM `ut_invitation_api_data` WHERE `id` = @reference_for_update);\n\t\n# Role in this unit for the invited user:\n\t#\t- Tenant 1\n\t# \t- Landlord 2\n\t#\t- Agent 5\n\t#\t- Contractor 3\n\t#\t- Management company 4\n\tSET @id_role_type = (SELECT `user_role_type_id` FROM `ut_invitation_api_data` WHERE `id` = @reference_for_update);\n\tSET @user_role_type_name = (SELECT `role_type` FROM `ut_role_types` WHERE `id_role_type` = @id_role_type);\n\t\n# The MEFE information:\n\tSET @mefe_invitor_user_id = (SELECT `mefe_invitor_user_id` FROM `ut_invitation_api_data` WHERE `id` = @reference_for_update);\n\t\n# What type of invitation is this?\n\tSET @invitation_type = (SELECT `invitation_type` FROM `ut_invitation_api_data` WHERE `id` = @reference_for_update);\n\t\t\t\t\t\t\t\n#################################################################\n#\n# All the variables have been set - we can call the procedures\n#\n#################################################################\n\n# Finalize the invitation\n\tCALL `finalize_invitation_to_a_case`;\n\n# Update the table 'ut_invitation_api_data' so we record what we have done\n\t\t\n\t# We do the update to record that we have reached the end of the script...\n\t\tUPDATE `ut_invitation_api_data`\n\t\t\tSET `api_post_datetime` = @mefe_invitation_sent\n\t\t\t\t, `script` = @this_script\n\t\t\tWHERE `mefe_invitation_id_int_value` = @mefe_invitation_id_int_value\n\t\t\t;\n" +
	"", etag: `"EE8ARl6v0wU="`})
