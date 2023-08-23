/**
 * OPERA Cloud Sales Event Management API
 * APIs to cater for Event Management functionality in OPERA Cloud. <br /><br />The Events feature in OPERA Cloud is designed to manage any kind of catering activity. Events can be as simple as a one-hour reception or more complex, such as a three-day business meeting with meals, breaks, and specific meeting functionSpaceDetails with setupCode and resource requirements. Any group function can be an Event.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 21.5.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
import type { BlockProfileTypeType } from './BlockProfileTypeType';
import type { ProfileType } from './ProfileType';
import type { UniqueIDType } from './UniqueIDType';
/**
 * A collection of Profile objects or Unique IDs of Profiles.
 * @export
 * @interface BlockProfilesType
 */
export interface BlockProfilesType {
    /**
     * Unique Id that references an object uniquely in the system.
     * @type {Array<UniqueIDType>}
     * @memberof BlockProfilesType
     */
    profileIdList?: Array<UniqueIDType>;
    /**
     *
     * @type {ProfileType}
     * @memberof BlockProfilesType
     */
    profile?: ProfileType;
    /**
     * Is this profile attached to block is primary?
     * @type {boolean}
     * @memberof BlockProfilesType
     */
    primary?: boolean;
    /**
     *
     * @type {BlockProfileTypeType}
     * @memberof BlockProfilesType
     */
    blockProfileType?: BlockProfileTypeType;
}
/**
 * Check if a given object implements the BlockProfilesType interface.
 */
export declare function instanceOfBlockProfilesType(value: object): boolean;
export declare function BlockProfilesTypeFromJSON(json: any): BlockProfilesType;
export declare function BlockProfilesTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): BlockProfilesType;
export declare function BlockProfilesTypeToJSON(value?: BlockProfilesType | null): any;
