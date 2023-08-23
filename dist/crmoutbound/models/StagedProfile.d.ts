/**
 * OPERA Cloud Customer Relationship Management Outbound API
 * APIs to cater for Customer Relationship Management external (outbound) functionality with OPERA. These APIs facilitate various operations related to getting data from an external system, and inserting it into OPERA.<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 22.3.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
import type { InstanceLink } from './InstanceLink';
import type { StagedProfileChangeInstructionsType } from './StagedProfileChangeInstructionsType';
import type { StagedProfileType } from './StagedProfileType';
/**
 * Request object for changing the staged profile.
 * @export
 * @interface StagedProfile
 */
export interface StagedProfile {
    /**
     *
     * @type {StagedProfileType}
     * @memberof StagedProfile
     */
    stagedProfile?: StagedProfileType;
    /**
     *
     * @type {StagedProfileChangeInstructionsType}
     * @memberof StagedProfile
     */
    changeInstructions?: StagedProfileChangeInstructionsType;
    /**
     *
     * @type {Array<InstanceLink>}
     * @memberof StagedProfile
     */
    links?: Array<InstanceLink>;
}
/**
 * Check if a given object implements the StagedProfile interface.
 */
export declare function instanceOfStagedProfile(value: object): boolean;
export declare function StagedProfileFromJSON(json: any): StagedProfile;
export declare function StagedProfileFromJSONTyped(json: any, ignoreDiscriminator: boolean): StagedProfile;
export declare function StagedProfileToJSON(value?: StagedProfile | null): any;
