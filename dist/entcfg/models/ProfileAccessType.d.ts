/**
 * OPERA Cloud Enterprise Configuration API
 * APIs to cater for Enterprise Configuration functionality in OPERA Cloud. <br /><br In this module, you can configure a variety of options related to your properties such as their locations, facilities, and local attractions. The available options are dependant on the active controls at your property.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 21.5.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
import type { ProfileSharedLevelType } from './ProfileSharedLevelType';
/**
 *
 * @export
 * @interface ProfileAccessType
 */
export interface ProfileAccessType {
    /**
     * Indicates the Chain code where the profile was created
     * @type {string}
     * @memberof ProfileAccessType
     */
    chainCode?: string;
    /**
     * Indicates the CRO code where the profile was created
     * @type {string}
     * @memberof ProfileAccessType
     */
    croCode?: string;
    /**
     * Indicates the Hotel code where the profile was created
     * @type {string}
     * @memberof ProfileAccessType
     */
    hotelId?: string;
    /**
     *
     * @type {ProfileSharedLevelType}
     * @memberof ProfileAccessType
     */
    sharedLevel?: ProfileSharedLevelType;
}
/**
 * Check if a given object implements the ProfileAccessType interface.
 */
export declare function instanceOfProfileAccessType(value: object): boolean;
export declare function ProfileAccessTypeFromJSON(json: any): ProfileAccessType;
export declare function ProfileAccessTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): ProfileAccessType;
export declare function ProfileAccessTypeToJSON(value?: ProfileAccessType | null): any;
